import { gql, useMutation } from "@apollo/client";
import Alert from "@mui/material/Alert";
import Typography from "@mui/material/Typography";
import { GoogleLogin } from "@react-oauth/google";
import { AppContext } from "app/index";
import { Mutation, MutationLoginArgs } from "graphql/types";
import React, { useContext, useState } from "react";

const LOGIN = gql`
  mutation login($credential: String!) {
    login(credential: $credential) {
      user {
        username
      }
      queryToken
    }
  }
`;

const Login = () => {
  const { setUser } = useContext(AppContext);
  const [googleError, setGoogleError] = useState(false);
  const [login, { error }] = useMutation<Mutation, MutationLoginArgs>(LOGIN);

  return (
    <>
      <Typography>Please sign in:</Typography>
      {(googleError || error) && (
        <Alert severity="error">
          Error signing in. Please try again later.
        </Alert>
      )}
      <GoogleLogin
        onSuccess={(resp) => {
          login({ variables: { credential: resp.credential } }).then(
            (response) => {
              setUser(response.data.login.user);
              window.postMessage(
                {
                  type: "heatmap-creator-query-token",
                  token: response.data.login.queryToken,
                },
                "*"
              );
            }
          );
          setGoogleError(false);
        }}
        onError={() => {
          setGoogleError(true);
        }}
      />
    </>
  );
};

export default Login;
