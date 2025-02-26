import { createContext, useContext, useState, ReactNode } from "react";

interface AlertContextType {
  showAlert: (message: string) => void;
}

const AlertContext = createContext<AlertContextType | undefined>(undefined);

export function AlertProvider({ children }: { children: ReactNode }) {
  const [alertMessage, setAlertMessage] = useState("");
  const [alertClassName, setAlertClassName] = useState("hidden");

  const showAlert = (message: string) => {
    setAlertMessage(message);
    setAlertClassName("fadeIn");

    setTimeout(() => {
      setAlertClassName("hidden");
    }, 3000);
  };

  return (
    <AlertContext.Provider value={{ showAlert }}>
      <div className="container mx-auto mt-8 max-w-screen-lg">
        <div className={`alert ${alertClassName}`}>{alertMessage}</div>
        {children}
      </div>
    </AlertContext.Provider>
  );
}

export function useAlert() {
  const context = useContext(AlertContext);
  if (!context) {
    throw new Error("useAlert must be used within an AlertProvider");
  }
  return context;
}
