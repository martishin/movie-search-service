import { XMarkIcon } from "@heroicons/react/24/outline";
import { createContext, ReactNode, useContext, useState } from "react";

interface AlertContextType {
  showAlert: (message: string) => void;
}

const AlertContext = createContext<AlertContextType | undefined>(undefined);

export function AlertProvider({ children }: { children: ReactNode }) {
  const [alertMessage, setAlertMessage] = useState("");
  const [isVisible, setIsVisible] = useState(false);
  const [isFading, setIsFading] = useState(false);

  const showAlert = (message: string) => {
    setAlertMessage(message);
    setIsVisible(true);
    setIsFading(false);

    setTimeout(() => {
      setIsFading(true);
      setTimeout(() => {
        setIsVisible(false);
      }, 500);
    }, 5000);
  };

  return (
    <AlertContext.Provider value={{ showAlert }}>
      {isVisible && (
        <div
          className={`fixed top-3 left-1/2 z-50 flex max-w-[80%] min-w-[250px] -translate-x-1/2 items-center justify-between rounded-lg bg-red-500 px-5 py-3 text-white shadow-lg transition-opacity duration-500 md:min-w-[350px] ${
            isFading ? "opacity-0" : "opacity-100"
          }`}
        >
          <span className="text-sm font-medium">{alertMessage}</span>

          <button
            className="ml-4 text-white hover:text-gray-200"
            onClick={() => setIsVisible(false)}
            aria-label="Close alert"
          >
            <XMarkIcon className="h-5 w-5" />
          </button>
        </div>
      )}

      {children}
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
