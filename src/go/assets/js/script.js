$gameTable = $('#gameTable');
$displaySpeed = $('#displaySpeed');
defaultDisplayTime = 250;

PageData = {
    rows: 10,
    cols: 12,
    startingTile: null,
    endingTile: null,
    pickingStartingTile: true,
}
GameMap = {
    rows: PageData.rows,
    cols: PageData.cols,
    // example tiles = [
    // [{id: 1, weight: 2, disabled: false}, {id: 2, weight: 4, disabled: true}],
    // ]
    tiles: [[]],
};

$(document).ready(function() {
    // generate map on page load
    generateMap(PageData.cols, PageData.rows)

    // update toastr options on load
    toastr.options.progressBar = true;
    toastr.options.closeButton = true;
});

function generateMap(cols, rows) {
    for (let i = 0; i < rows; i++) {
        GameMap.tiles[i] = [];
        for (let j = 0; j < cols; j++) {
            GameMap.tiles[i][j] = {
                id: (i * cols) + j + 1,
                weight: Math.floor(Math.random() * 10) + 1,
                disabled: false
            };
        }
    }
    drawMap(GameMap);
}

function drawMap(gameMap) {
    let map = '<tbody>';

    for (let i = 0; i < gameMap.rows; i++) {
        map += '<tr>';
        for (let j = 0; j < gameMap.cols; j++) {
            map += '<td><button ' +
                'type="button" ' +
                'class="btn btn-light shadow-none map-tile" ' +
                `id=${gameMap.tiles[i][j].id} ` +
                'onclick="tileClicked(this)">' +
                `${gameMap.tiles[i][j].weight}` +
                '</button></td>';
        }
        map += '</tr>';
    }
    map += '</tbody>';
    $gameTable.append(map);
}

function tileClicked(tile) {
    if (getTileById(tile.id).disabled) {
        return;
    }

    if ($('#flexSwitchCheckDefault').is(':checked')) {
        handleMountains(tile);
        return;
    }

    // switch between selecting Start and End tiles
    if (PageData.pickingStartingTile) {
        if (PageData.startingTile != null) {
            $('#' + PageData.startingTile)
                .removeClass('btn-secondary')
                .prop('disabled', false)
                .text(getTileById(PageData.startingTile).weight);
        }

        $(tile)
            .addClass('btn-secondary')
            .prop('disabled', true)
            .text("Start");

        PageData.startingTile = parseInt(tile.id);

        PageData.pickingStartingTile = !PageData.pickingStartingTile;
    } else {
        if (PageData.endingTile != null) {
            $('#' + PageData.endingTile)
                .removeClass('btn-primary')
                .prop('disabled', false)
                .text(getTileById(PageData.endingTile).weight);
        }

        $(tile)
            .addClass('btn-primary')
            .prop('disabled', true)
            .text("End");

        PageData.endingTile = parseInt(tile.id);

        PageData.pickingStartingTile = !PageData.pickingStartingTile;
    }
}

async function getPath(button) {
    button.disabled = true;

    let data = {
        start: PageData.startingTile,
        end: PageData.endingTile,
        gameMap: GameMap,
    };

    await fetch('/dijkstra', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    })
        .then(response => response.json())
        .then(async paths => {
            if (paths.error) {
                return Promise.reject(paths.error);
            }

            console.log('Successfully fetched paths');
            await drawPath(paths.explored, 'explored');
            await drawPath(paths.shortest, 'shortest');
        })
        .catch((error) => {
            console.error('Error: ', error);
            toastr.error(error);
        });

    button.disabled = false;
}

// drawPath is a helper function that draws the path on the map
async function drawPath(path, className) {
    let shortestPathCost = 0;
    let totalTiles = getExplorableTilesNum();

    for (let i = 0; i < path.length; i++) {
        // speed up based on user input
        await new Promise(r => setTimeout(r, defaultDisplayTime - (defaultDisplayTime * $displaySpeed.val() / 100)));

        // update mini-card metrics
        if (className === 'explored') {
            // explored map coverage mini-card
            let exploredPercentage = Math.floor((i / totalTiles) * 100);
            $('#coverageText').text(`${exploredPercentage} %`);
            $('#coverageBar').css('width', `${exploredPercentage}%`);

            // tiles explored mini-card
            $('#tilesExploredCount').text(`${i}`);

            // efficiency rating mini-card
            let efficiencyRating = 100 - exploredPercentage;
            $('#efficiencyText').text(`${efficiencyRating} %`);
            $('#efficiencyBar').css('width', `${efficiencyRating}%`);

        } else if (className === 'shortest') {
            // shortest path cost mini-card
            shortestPathCost += getTileById(path[i]).weight;
            $('#shortestPathCost').text(`${shortestPathCost}`);
        }

        // draw path
        $('#' + path[i])
            .addClass(className)
            .prop('disabled', true)
    }
}

// handleMountains is a helper function that draws the mountains feature
function handleMountains(tile) {
    if (parseInt(tile.id) === PageData.startingTile || parseInt(tile.id) === PageData.endingTile) {
        return;
    }
    getTileById(tile.id).disabled = true;
    $(tile)
        .removeClass()
        .addClass('btn btn-dark map-tile');
}

// clearMap resets the map to its original state
function clearMap() {
    $('.map-tile').each(function(_, tile) {
        // replace all tiles with default settings
        getTileById(tile.id).disabled = false;
        $(tile)
            .removeClass()
            .addClass('btn btn-light shadow-none map-tile')
            .prop('disabled', false)
    });
}

// getTileById gets the tile from GameMap with the id
function getTileById(tileId) {
    if (typeof tileId === 'string' || tileId instanceof String) {
        tileId = parseInt(tileId);
    }

    let row = Math.floor((tileId - 1) / PageData.cols);
    let col = (tileId - 1) % PageData.cols;
    return GameMap.tiles[row][col];
}

// getExplorableTilesNum gets the number of explorable tiles
function getExplorableTilesNum() {
    let explorableTiles = PageData.rows * PageData.cols;
    for (let i = 0; i < explorableTiles; i++) {
        if (getTileById(i+1).disabled) {
            explorableTiles--;
        }
    }
    return explorableTiles;
}