import Link from "@mui/material/Link";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import React from "react";

export const Footer = () => {
  return (
    <Stack
      direction="column"
      spacing={1}
      sx={{
        bottom: 0,
        textAlign: "center",
        padding: "20px 0",
      }}
    >
      <Typography>Copyright ©{new Date().getFullYear()}</Typography>
      <Typography>
        <Link href="/privacy">Privacy Policy</Link>
      </Typography>
      <Typography>
        <Link href="mailto:help@getheatmap-creator.app">Contact Support</Link>
      </Typography>
    </Stack>
  );
};
