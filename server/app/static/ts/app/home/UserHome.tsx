import { gql, useQuery } from "@apollo/client";
import Typography from "@mui/material/Typography";
import { AppContext } from "app/index";
import { Query } from "graphql/types";
import React, { useContext } from "react";

const USER_STATS = gql`
  query getStats {
    getUserStats {
      maxQueries
      numQueries
    }
  }
`;

const UserHome = () => {
  const { user } = useContext(AppContext);
  const { data, loading, error } = useQuery<Query>(USER_STATS);

  if (loading) {
    return <Typography>Loading...</Typography>;
  }

  if (error) {
    return (
      <Typography>
        Error fetching user information. Please try again later.
      </Typography>
    );
  }

  return (
    <>
      <Typography>Logged in as: {user.username} </Typography>
      <Typography>
        Daily queries used: {data.getUserStats.numQueries}
      </Typography>
      <Typography>
        Daily queries allowed: {data.getUserStats.maxQueries}
      </Typography>
      <Typography fontStyle="italic" fontWeight="bold">
        Want more queries? Paid plans coming soon!
      </Typography>
    </>
  );
};

export default UserHome;
