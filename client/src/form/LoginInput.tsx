import React, { forwardRef } from "react";

interface LoginInputProps {
  title: string;
  name: string;
  type: string;
  autoComplete: string;
  onChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
  hasError?: boolean;
  errorMsg?: string;
}

const LoginInput = forwardRef<HTMLInputElement, LoginInputProps>((props, ref) => {
  const { title, name, type, autoComplete, onChange, hasError, errorMsg } = props;
  const errorId = `${name}-error`;

  return (
    <div className="mt-3">
      <label
        htmlFor={name}
        className="block text-center text-sm leading-6 font-medium text-gray-900"
      >
        {title}
      </label>
      <div className="mt-2">
        <input
          id={name}
          name={name}
          type={type}
          autoComplete={autoComplete}
          className={`block w-full rounded-md border-0 px-3 py-1.5 text-gray-900 ring-1 ring-inset ${
            hasError ? "ring-red-600" : "ring-gray-300"
          } placeholder:text-gray-400 focus:ring-2 focus:ring-blue-600 focus:ring-inset`}
          ref={ref}
          onChange={onChange}
          aria-invalid={hasError ? "true" : "false"}
          aria-describedby={hasError ? errorId : undefined}
        />
        {hasError && (
          <p id={errorId} className="mt-2 text-sm text-red-600">
            {errorMsg}
          </p>
        )}
      </div>
    </div>
  );
});

LoginInput.displayName = "LoginInput";

export default LoginInput;
