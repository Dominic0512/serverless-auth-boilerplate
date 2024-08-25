import { Outlet } from "@remix-run/react";

export default function Auth() {
  return (
    <div className="font-sans p-4">
      <h1 className="text-3xl">Auth</h1>
      <Outlet />
    </div>
  );
}