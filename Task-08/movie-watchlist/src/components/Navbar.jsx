import { Link } from "react-router-dom";

export default function Navbar() {
  return (
    <nav className="navbar">
      <Link to="/">🎬 MovieWatchlist</Link>
      <Link to="/watchlist">My Watchlist</Link>
    </nav>
  );
}