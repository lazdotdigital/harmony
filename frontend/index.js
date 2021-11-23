const renderSongTree = () => {
  const songTree = document.querySelector(".song-tree");
  songTree.innerHTML = "";

  go_getSongs().then((songs) => {
    const artistToSongsObj = {};
    songs.forEach(({ title, by, audioPath }) => {
      if (!artistToSongsObj[by]) artistToSongsObj[by] = [];
      artistToSongsObj[by].push({ title, audioPath });
    });

    Object.keys(artistToSongsObj).forEach((artist) => {
      const span = document.createElement("span");
      span.className = "song-tree__caret";

      const node = document.createTextNode(artist);
      span.appendChild(node);

      const li = document.createElement("li");
      li.className = "song-tree__artist";
      li.appendChild(span);
      songTree.appendChild(li);
    });

    const player = document.querySelector(".app__player");
    const artistLis = document.querySelectorAll(".song-tree__artist");
    Object.values(artistToSongsObj).forEach((songs, i) => {
      const ul = document.createElement("ul");
      ul.className = "song-tree__songs";

      songs.forEach((song) => {
        const li = document.createElement("li");
        li.addEventListener("click", () => {
          player.src = song.audioPath;
        });

        const node = document.createTextNode(song.title);
        li.appendChild(node);
        ul.appendChild(li);
      });

      artistLis[i].appendChild(ul);
    });

    document.querySelectorAll(".song-tree__caret").forEach((el) => {
      el.addEventListener("click", () => {
        el.parentElement
          .querySelector(".song-tree__songs")
          .classList.toggle("song-tree__songs--active");

        el.classList.toggle("song-tree__caret--down");
      });
    });
  });
};

renderSongTree();

const addSongForm = document.querySelector(".app__add-song");
addSongForm.addEventListener("submit", (e) => {
  e.preventDefault();

  const { children } = addSongForm;
  const url = children[0].value;
  const title = children[1].value;
  const by = children[2].value;

  go_getAudio(url)
    .then((audio) => {
      go_addSong(title, by, audio)
        .then(() => {
          renderSongTree();
        })
        .catch((err) => {
          console.error(err);
        });
    })
    .catch((err) => {
      console.error(err);
    });
});
