function solution(genres, plays) {
  let answer = [];
  const totalCount = {};
  const songs = [];

  for (let i = 0; i < genres.length; i++) {
    totalCount[genres[i]] = (totalCount[genres[i]] || 0) + plays[i];
    songs[i] = [genres[i], plays[i], i];
  }
  answer = Object.entries(totalCount).sort((a, b) => b[1] - a[1]);
  answer.map((count, i) => {
    songs.filter((data) => {
      data[0] === count[0];
    });
  });

  return answer
    .map((count, i) =>
      songs
        .filter((data) => data[0] === count[0])
        .sort((a, b) => b[1] - a[1])
        .map((data) => data[2])
        .slice(0, 2)
    )
    .flat();
}