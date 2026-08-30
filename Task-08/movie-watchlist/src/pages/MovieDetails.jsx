import { useParams, useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";
import { fetchMovieDetails, imgUrl } from "../api/tmdb";
import { useWatchlist } from "../context/WatchlistContext";

export default function MovieDetails() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [movie, setMovie] = useState(null);
  const { addToWatchlist, removeFromWatchlist, isInWatchlist } = useWatchlist();

  useEffect(() => {
    fetchMovieDetails(id).then(setMovie);
  }, [id]);

  if (!movie) return <p>Loading...</p>;
  const saved = isInWatchlist(movie.id);

  return (
    <div className="movie-details">
      <button onClick={() => navigate(-1)}>← Back</button>
      <img src={imgUrl(movie.poster_path)} alt={movie.title} />
      <h1>{movie.title} ({movie.release_date?.slice(0, 4)})</h1>
      <p><strong>Rating:</strong> {movie.vote_average?.toFixed(1)} / 10</p>
      <p><strong>Genres:</strong> {movie.genres?.map((g) => g.name).join(", ")}</p>
      <p>{movie.overview}</p>
      <button onClick={() => (saved ? removeFromWatchlist(movie.id) : addToWatchlist(movie))}>
        {saved ? "Remove from Watchlist" : "Add to Watchlist"}
      </button>
    </div>
  );
}