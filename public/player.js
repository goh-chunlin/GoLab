function watchVideoGeneral(videoID) {

        document.getElementById('ytplayer').src = "https://www.youtube.com/embed/" + videoID + "?autoplay=1&loop=1&playlist=" + videoID;

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

function updateVideoInfo(videoId, videoName) {
        $('#modalUpdateVideoInfo').modal('show');

        $('#modalUpdateVideoInfoVideoId').val(videoId);
        $('#modalUpdateVideoInfoName').val(videoName);
}

function deleteVideo(videoId, videoName) {
        $('#modalDeleteVideo').modal('show');

        $('#modalDeleteVideoVideoId').val(videoId);
        $('#modalDeleteVideoName').html(videoName);
}