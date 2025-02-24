import { JSX } from "react";

interface PageHeaderProps {
  title: string;
}

export default function PageHeader({ title }: PageHeaderProps): JSX.Element {
  return (
    <div>
      <h2 className="text-center text-xl font-bold tracking-tight">{title}</h2>
      <hr className="mt-1 border-gray-300" />
    </div>
  );
}
