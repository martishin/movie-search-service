import { JSX } from "react";
import PageHeader from "../components/layout/PageHeader";
import { Link } from "react-router";
import ticketLogo from "../assets/movie_tickets.jpg";

export default function Home(): JSX.Element {
  return (
    <div className="text-center">
      <PageHeader title="Find a movie to watch tonight!" />
      <Link to="/movies">
        <img src={ticketLogo} alt="movie tickets" className="m-auto w-40 pt-3"></img>
      </Link>
    </div>
  );
}
