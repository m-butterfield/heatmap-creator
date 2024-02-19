import Link from "@mui/material/Link";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import Login from "app/home/Login";
import UserHome from "app/home/UserHome";
import { AppContext } from "app/index";
import React, { useContext } from "react";

const Home = () => {
  const { user } = useContext(AppContext);

  return (
    <Stack direction="column" spacing={3} alignItems="center">
      <Typography variant="h3">heatmap-creator</Typography>
      <Link
        href="https://chrome.google.com/webstore/detail/heatmap-creator/knpghemhjolccmkhdpihiajlpbknmfpa"
        target="_blank"
      >
        Install heatmap-creator
      </Link>
      {user ? <UserHome /> : <Login />}
    </Stack>
  );
};

export default Home;
