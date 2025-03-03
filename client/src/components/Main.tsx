import { createBrowserRouter, RouterProvider } from "react-router";
import App from "./App";
import ErrorPage from "../pages/ErrorPage";
import Home from "../pages/Home";
import { AlertProvider } from "../context/AlertContext";
import { AuthProvider } from "../context/AuthContext";
import Login from "../pages/Login";
import SignUp from "../pages/SignUp";
import Movies from "../pages/Movies";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      { index: true, element: <Home /> },
      { path: "/movies", element: <Movies /> },
      { path: "/login", element: <Login /> },
      { path: "/sign-up", element: <SignUp /> },
    ],
  },
]);

export default function Main() {
  return (
    <AuthProvider>
      <AlertProvider>
        <div className="min-h-screen bg-[#F4F4F4]">
          <RouterProvider router={router} />
        </div>
      </AlertProvider>
    </AuthProvider>
  );
}
