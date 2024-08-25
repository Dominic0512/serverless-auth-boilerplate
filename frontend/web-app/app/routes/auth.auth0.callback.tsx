import { useLocation } from "@remix-run/react";

export default function Auth0Callback() {
  const { search } = useLocation();
  return (
    <div className="font-sans p-4">
      <h1 className="text-3xl">Auth0Callback</h1>
      <div>{search}</div>
    </div>
  );
}
