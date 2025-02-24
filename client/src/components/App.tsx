import { useEffect, useState, useCallback, JSX } from "react";
import UserDetails from "../models/UserDetails";
import { Outlet, useNavigate } from "react-router";
import Navigation from "./layout/Navigation";
import Header from "./layout/Header";
import Alert from "./Alert";

export default function App(): JSX.Element | null {
  const navigate = useNavigate();

  const [userDetails, setUserDetails] = useState<UserDetails | null>(null);
  const [isFetchingAuth, setIsFetchingAuth] = useState(true);
  const [alertMessage, setAlertMessage] = useState("");
  const [alertClassName, setAlertClassName] = useState("hidden");

  const fetchUserDetails = useCallback(async () => {
    try {
      const res = await fetch("/api/user", { credentials: "include" });

      if (!res.ok) {
        setUserDetails(null);
        setIsFetchingAuth(false);
        return;
      }

      const userData = await res.json();
      setUserDetails(userData);
    } catch (err) {
      console.error("Error fetching user:", err);
      navigate("/");
    } finally {
      setIsFetchingAuth(false);
    }
  }, [navigate]);

  useEffect(() => {
    fetchUserDetails();
  }, [fetchUserDetails]);

  useEffect(() => {
    if (alertClassName === "fadeIn") {
      const timer = setTimeout(() => {
        setAlertClassName("hidden");
      }, 3000);

      return () => clearTimeout(timer);
    }
  }, [alertClassName]);

  if (isFetchingAuth) {
    return null;
  }

  return (
    <div className="container mx-auto mt-8 max-w-screen-lg">
      <Alert message={alertMessage} alertClassName={alertClassName} />
      <Header userDetails={userDetails} setUserDetails={setUserDetails} />
      <div className="mt-4 flex">
        <div className="w-48">
          <Navigation userDetails={userDetails} />
        </div>
        <div className="ml-4 mr-4 w-min flex-grow">
          <Outlet
            context={{
              userDetails,
              setUserDetails,
              setAlertMessage,
              setAlertClassName,
            }}
          />
        </div>
      </div>
    </div>
  );
}
