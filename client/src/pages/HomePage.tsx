import { JSX } from "react";
import { Link } from "react-router";
import ticketLogo from "../assets/movie_tickets.jpg";

export default function HomePage(): JSX.Element {
  return (
    <div className="text-center">
      <Link to="/movies">
        <img src={ticketLogo} alt="movie tickets" className="m-auto w-40 pt-3"></img>
      </Link>
    </div>
  );
}
