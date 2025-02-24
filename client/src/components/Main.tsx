import { createBrowserRouter, RouterProvider } from "react-router";
import App from "./App";
import ErrorPage from "./pages/ErrorPage";
import Home from "./pages/Home";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [{ index: true, element: <Home /> }],
  },
]);

export default function Main() {
  return <RouterProvider router={router} />;
}
