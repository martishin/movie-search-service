import { XCircleIcon } from "@heroicons/react/20/solid";
import { JSX } from "react";

interface AlertProps {
  alertClassName: string;
  message: string;
}

export default function Alert({ alertClassName, message }: AlertProps): JSX.Element {
  return (
    <div
      className={`alert ${alertClassName} fixed top-3 left-1/2 w-3/5 max-w-2xl -translate-x-1/2 transform`}
    >
      <div className="rounded-md bg-red-50 p-4">
        <div className="flex">
          <div className="flex-shrink-0">
            <XCircleIcon className="h-5 w-5 text-red-400" aria-hidden="true" />
          </div>
          <div className="ml-3">
            <p className="text-sm font-medium text-red-800">{message}</p>
          </div>
        </div>
      </div>
    </div>
  );
}
