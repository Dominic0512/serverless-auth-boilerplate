import type { MetaFunction } from "@remix-run/node";
import { useAuth0 } from "@auth0/auth0-react";


export const meta: MetaFunction = () => {
  return [
    { title: "Serverless Auth Boilerplate" },
    { name: "description", content: "This is Serverless Auth Boilerplate." },
  ];
};

export default function Index() {
  const { loginWithRedirect } = useAuth0();
  return (
    <div className="font-sans p-4">
      <div className="w-100 h-100 flex flex-col items-center justify-center gap-4">
        <div>
          <h1 className="text-3xl">Welcome to Serverless Auth Boilerplate</h1>
        </div>
        <div>
          <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded" onClick={() => loginWithRedirect() }>Login By Auth0</button>
        </div>     
      </div>  
    </div>
  );
}
