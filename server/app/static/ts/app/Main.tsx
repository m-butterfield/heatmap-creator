import { useQuery } from "@apollo/client";
import Box from "@mui/material/Box";
import Container from "@mui/material/Container";
import Typography from "@mui/material/Typography";
import AppRoutes from "app/AppRoutes";
import { Footer } from "app/Footer";
import { Header } from "app/Header";
import { AppContext } from "app/index";
import { ME } from "app/queries";
import { Query } from "graphql/types";
import React, { useContext, useEffect } from "react";
import { BrowserRouter } from "react-router-dom";

const Main = () => {
  const { setUser } = useContext(AppContext);
  const { data, loading, error } = useQuery<Query>(ME);
  useEffect(() => {
    if (data && data.me) {
      setUser(data.me);
    }
  }, [data]);
  return (
    <>
      {loading || error ? (
        <Typography>{error && "something is wrong"}</Typography>
      ) : (
        <BrowserRouter>
          <Header />
          <Container>
            <Box sx={{ my: 2, minHeight: "100vh" }}>
              <AppRoutes />
            </Box>
          </Container>
          <Footer />
        </BrowserRouter>
      )}
    </>
  );
};

export default Main;
