import MovieGrid from "../components/MovieGrid";
import { useWatchlist } from "../context/WatchlistContext";

export default function Watchlist() {
  const { watchlist } = useWatchlist();
  return (
    <div>
      <h2>My Watchlist ({watchlist.length})</h2>
      <MovieGrid movies={watchlist} />
    </div>
  );
}