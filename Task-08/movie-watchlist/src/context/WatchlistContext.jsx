import { createContext, useContext, useState, useEffect } from "react";

const WatchlistContext = createContext();

export function WatchlistProvider({ children }) {
  const [watchlist, setWatchlist] = useState(() => {
    const saved = localStorage.getItem("watchlist");
    return saved ? JSON.parse(saved) : [];
  });

  useEffect(() => {
    localStorage.setItem("watchlist", JSON.stringify(watchlist));
  }, [watchlist]);

  const addToWatchlist = (movie) =>
    setWatchlist((prev) =>
      prev.some((m) => m.id === movie.id) ? prev : [...prev, movie]
    );

  const removeFromWatchlist = (id) =>
    setWatchlist((prev) => prev.filter((m) => m.id !== id));

  const isInWatchlist = (id) => watchlist.some((m) => m.id === id);

  return (
    <WatchlistContext.Provider
      value={{ watchlist, addToWatchlist, removeFromWatchlist, isInWatchlist }}
    >
      {children}
    </WatchlistContext.Provider>
  );
}

export const useWatchlist = () => useContext(WatchlistContext);