import { ApolloClient, ApolloProvider, InMemoryCache } from "@apollo/client";
import CssBaseline from "@mui/material/CssBaseline";
import ThemeProvider from "@mui/material/styles/ThemeProvider";
import { GoogleOAuthProvider } from "@react-oauth/google";
import { googleOAuthClientID } from "app/constants";
import Main from "app/Main";
import { theme } from "app/theme";
import { User } from "graphql/types";
import React, { createContext, useState } from "react";
import { createRoot } from "react-dom/client";

type AppState = {
  user: User | null;
  setUser: (user: User) => void;
};

export const AppContext = createContext<AppState>(null);

const client = new ApolloClient({
  uri: "/graphql/query",
  cache: new InMemoryCache(),
  credentials: "same-origin",
});

const App = () => {
  const [stateUser, setUser] = useState<User | null>(null);
  const state = {
    user: stateUser,
    setUser: setUser,
  };
  return (
    <AppContext.Provider value={state}>
      <ApolloProvider client={client}>
        <ThemeProvider theme={theme}>
          <GoogleOAuthProvider clientId={googleOAuthClientID}>
            <CssBaseline />
            <Main />
          </GoogleOAuthProvider>
        </ThemeProvider>
      </ApolloProvider>
    </AppContext.Provider>
  );
};

const container = document.getElementById("root");
const root = createRoot(container);

root.render(<App />);
