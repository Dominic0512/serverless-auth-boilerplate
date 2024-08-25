import {
  json,
  Links,
  LiveReload,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
  useLoaderData,
} from "@remix-run/react";
import { Auth0Provider } from "@auth0/auth0-react";
import "./tailwind.css";

export async function loader() {
  return json({
    ENV: {
      AUTH0_DOMAIN: process.env.AUTH0_DOMAIN,
      AUTH0_CLIENT_ID: process.env.AUTH0_CLIENT_ID,
      AUTH0_REDIRECT_URI: process.env.AUTH0_REDIRECT_URI,
    }
  })
}

export function Layout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <Meta />
        <Links />
      </head>
      <body>
        {children}
        <ScrollRestoration />
        <Scripts />
        <LiveReload />
      </body>
    </html>
  );
}

export default function App() {
  const data = useLoaderData<typeof loader>();
  return (
    <Auth0Provider 
      domain={data.ENV.AUTH0_DOMAIN || ""}
      clientId={data.ENV.AUTH0_CLIENT_ID || ""}
      authorizationParams={{ redirect_uri: data.ENV.AUTH0_REDIRECT_URI || "" }}
    >
      <Outlet />
    </Auth0Provider>
  
  );
}
