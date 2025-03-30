import { createBrowserRouter, RouterProvider } from "react-router";

import { AlertProvider } from "../context/AlertContext";
import { AuthProvider } from "../context/AuthContext";
import App from "./App";
import ErrorPage from "./ErrorPage";
import LoginPage from "./LoginPage";
import MoviePage from "./MoviePage";
import MoviesPage from "./MoviesPage";
import SignUpPage from "./SignUpPage";
import WatchLikedPage from "./WatchLikedPage";
import WatchOnlinePage from "./WatchOnlinePage";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      { index: true, element: <WatchOnlinePage /> },
      { path: "/login", element: <LoginPage /> },
      { path: "/sign-up", element: <SignUpPage /> },
      { path: "/movies", element: <MoviesPage /> },
      { path: "/movies/:id", element: <MoviePage /> },
      { path: "/liked", element: <WatchLikedPage /> },
    ],
  },
]);

export default function Main() {
  return (
    <AuthProvider>
      <AlertProvider>
        <div className="bg-[#F4F4F4]">
          <RouterProvider router={router} />
        </div>
      </AlertProvider>
    </AuthProvider>
  );
}
