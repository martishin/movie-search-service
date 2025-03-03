import Genre from "../models/Genre";

type GenreTagProps = {
  genre: Genre;
};

const genreColors = [
  "bg-blue-200 text-blue-800",
  "bg-green-200 text-green-800",
  "bg-yellow-200 text-yellow-800",
  "bg-red-200 text-red-800",
  "bg-purple-200 text-purple-800",
  "bg-pink-200 text-pink-800",
  "bg-indigo-200 text-indigo-800",
  "bg-orange-200 text-orange-800",
  "bg-teal-200 text-teal-800",
  "bg-lime-200 text-lime-800",
  "bg-cyan-200 text-cyan-800",
  "bg-amber-200 text-amber-800",
  "bg-emerald-200 text-emerald-800",
  "bg-violet-200 text-violet-800",
  "bg-rose-200 text-rose-800",
  "bg-sky-200 text-sky-800",
  "bg-fuchsia-200 text-fuchsia-800",
  "bg-gray-200 text-gray-800",
  "bg-stone-200 text-stone-800",
  "bg-zinc-200 text-zinc-800",
  "bg-neutral-200 text-neutral-800",
];

const getGenreColor = (genreId: number) => {
  return genreColors[genreId % genreColors.length];
};

export default function GenreTag({ genre }: GenreTagProps) {
  return (
    <span className={`rounded-full px-2 py-1 text-xs font-medium ${getGenreColor(genre.id)}`}>
      {genre.genre}
    </span>
  );
}
