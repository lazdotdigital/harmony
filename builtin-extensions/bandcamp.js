exports.Name = "Bandcamp";

exports.Claim = function (url) {
  // TODO: make this into a regex?
  var parts = url.toLowerCase().split(".");
  if (parts.length !== 3 || parts[1] !== "bandcamp") return false;

  parts = parts[2].split("/");
  if (parts.length !== 3 || parts[1] !== "track") return false;

  return true;
};

exports.GetAudio = function (url) {
  var html = fetchString(url).toString();
  var idx = html.search("https://t4.bcbits.com/stream/");
  html = html.slice(idx);
  html = html.slice(0, html.search("&quot;"));

  return fetchBytes(html);
};
