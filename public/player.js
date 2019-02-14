function watchVideoGeneral(videoID) {

    document.getElementById('ytplayer').src = "https://www.youtube.com/embed/" + videoID + "?autoplay=1&loop=1&playlist=" + videoID;

}

function loadVideos() {
    $.ajax(
        {
            url: "/api/video/",
            cache: false,
            success: function (result) {
                availableVideos = result;
                vmVideos.videos = availableVideos;
                
                $('#storedVideosList').slimScroll(
                    {
                        height: '500px',
                        size: '8px',
                        alwaysVisible: true,
                        railVisible: true
                    }
                );
            },
            error: function (e) {

                addSystemMessage('alert-danger', 'HTTP ' + e.status + ': ' + e.statusText + '.');

            }
        }
    );
}

function addVideo() {
    $.ajax(
        {
            url: "/api/video/",
            cache: false,
            type: "POST",
            data: JSON.stringify({
                videoTitle: $('#hidVideoName').val(),
                url: 'https://www.youtube.com/watch?v=' + $('#hidVideoID').val()
            }),
            dataType: "json",
            success: function (result) {

                if (result.status) {
                    loadVideos();
                    addSystemMessage('alert-success', result.message);
                } else {
                    addSystemMessage('alert-danger', result.message);
                }

            },
            error: function (e) {

                addSystemMessage('alert-danger', 'HTTP ' + e.status + ': ' + e.statusText + '.');

            },
            complete: function (e) {

                $('#hidVideoID').val('');
                $('#hidVideoName').val('');
                $('#btnAddToList').attr("disabled", true);

            }
        }
    );
}

function updateVideo() {
    $.ajax(
        {
            url: "/api/video/" + $('#modalUpdateVideoInfoVideoId').val(),
            cache: false,
            type: "PUT",
            data: JSON.stringify({                
                videoTitle: $('#modalUpdateVideoInfoName').val()
            }),
            dataType: "json",
            success: function (result) {

                if (result.status) {
                    loadVideos();
                    addSystemMessage('alert-success', result.message);
                } else {
                    addSystemMessage('alert-danger', result.message);
                }

            },
            error: function (e) {

                addSystemMessage('alert-danger', 'HTTP ' + e.status + ': ' + e.statusText + '.');

            },
            complete: function (e) {

                $('#modalUpdateVideoInfo').modal('hide');

            }
        }
    );
}

function deleteVideo() {
    console.log('Deleting');
    $.ajax(
        {
            url: "/api/video/" + $('#modalDeleteVideoVideoId').val(),
            cache: false,
            type: "DELETE",
            data: JSON.stringify({}),
            dataType: "json",
            success: function (result) {

                if (result.status) {
                    loadVideos();
                    addSystemMessage('alert-success', result.message);
                } else {
                    addSystemMessage('alert-danger', result.message);
                }

            },
            error: function (e) {

                addSystemMessage('alert-danger', 'HTTP ' + e.status + ': ' + e.statusText + '.');

            },
            complete: function (e) {

                $('#modalDeleteVideo').modal('hide');

            }
        }
    );
}

function watchVideo() {
    var videoURL = $('#txtVideoURL').val();
    $('#hidVideoID').val('');
    $('#hidVideoName').val('');
    $('#btnAddToList').attr("disabled", true);

    var match,
        pl = /\+/g,  // Regex for replacing addition symbol with a space
        search = /([^&=]+)=?([^&]*)/g,
        decode = function (s) { return decodeURIComponent(s.replace(pl, " ")); };

    var urlParams = {};
    while (match = search.exec(videoURL.slice(videoURL.indexOf('?') + 1))) {
        urlParams[decode(match[1])] = decode(match[2]);
    }

    var videoID = urlParams['v'];
    if (typeof urlParams['v'] === 'undefined') {
        alert('Please enter a valid YouTube video URL in the textbox field. For example, https://www.youtube.com/watch?v=0tt4TeFcFFI.');
    } else {
        watchVideoGeneral(videoID);

        $.getJSON('https://www.googleapis.com/youtube/v3/videos?part=snippet,contentDetails&id=' + videoID + '&key=AIzaSyD2phoI-AKDI26thimy-FYe9EISDFScIu0',
            function () { })
            .done(
                function (data) {
                    var videoInfoItem = data.items[0];
                    var videoTitle = ((videoInfoItem.snippet.title).replace(/"/g, '&quot;')).replace(/'/g, '&apos;');

                    $('#hidVideoID').val(videoID);
                    $('#hidVideoName').val(videoTitle);
                    $('#txtVideoURL').val('');
                    $('#btnAddToList').attr("disabled", false);

                    $('#storedVideosList').slimScroll({
                        height: '500px',
                        size: '8px',
                        alwaysVisible: true,
                        railVisible: true
                    });
                }
            )
            .fail(
                function () {
                    $('#hidVideoID').val('');
                    $('#hidVideoName').val('');
                    $('#txtVideoURL').val('');
                }
            );
    }
}

function showUpdateVideoInfoPopup(videoId, videoName) {
    $('#modalUpdateVideoInfo').modal('show');

    $('#modalUpdateVideoInfoVideoId').val(videoId);
    $('#modalUpdateVideoInfoName').val(videoName);
}

function showDeleteVideoPopup(videoId, videoName) {
    $('#modalDeleteVideo').modal('show');

    $('#modalDeleteVideoVideoId').val(videoId);
    $('#modalDeleteVideoName').html(videoName);
}

function addSystemMessage(alertClassName, message) {

    $('#systemMessagePanel').append(
        '<div class="alert alert-dismissible ' + alertClassName + '">' +
            '<button type="button" class="close" data-dismiss="alert">&times;</button>' +
            message +
        '</div>');

}