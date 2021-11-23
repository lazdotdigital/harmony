exports.Name = "YouTube";

function id(url) {
  var re =
    /(?:youtube\.com\/(?:[^\/]+\/.+\/|(?:v|e(?:mbed)?)\/|.*[?&]v=)|youtu\.be\/)([^"&?\/\s]{11})/gi;

  res = re.exec(url);
  if (res === null) return null;
  return res[1];
}

exports.Claim = function (url) {
  if (id(url) === null) return false;
  return true;
};

exports.GetAudio = function (url) {
  return youtubeToMP3(id(url));
};
