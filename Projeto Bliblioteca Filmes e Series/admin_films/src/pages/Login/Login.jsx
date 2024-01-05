import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import CssBaseline from "@mui/material/CssBaseline";
import Grid from "@mui/material/Grid";
import Paper from "@mui/material/Paper";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
import * as React from "react";
import { useState } from "react";
import toast from "react-hot-toast";
import { useNavigate } from "react-router-dom";
import { handleLogin } from "../../api/login";
import { CircularProgress } from "@mui/material";

const defaultTheme = createTheme();

export default function Login() {
  const [loading, setLoading] = useState(false);

  const navigate = useNavigate();

  const handleSubmit = async (event) => {
    event.preventDefault();
    const data = new FormData(event.currentTarget);
    setLoading(true);
    var result = await handleLogin(data.get("username"), data.get("password"));
    setLoading(false);

    if (result === true) {
      navigate("/catalogo");
    } else {
      toast.error("Houve um erro ao tentar se conectar, tente novamente.");
    }
  };

  return (
    <ThemeProvider theme={defaultTheme}>
      <Grid container component="main" sx={{ height: "100vh" }}>
        <CssBaseline />
        <Grid item xs={12} sm={8} md={12} component={Paper} elevation={6} square>
          <Box
            sx={{
              // my: 8,
              mx: 4,
              display: "flex",
              height: "100%",
              flexDirection: "column",
              alignItems: "center",
              justifyContent: "center",
            }}
          >
            <Box sx={{ textAlign: "center" }}>
              <Typography style={{ fontSize: "52px", fontWeight: "bolder" }}>
                CATALOGO
              </Typography>
              <h4 style={{ fontStyle: "italic" }}>Painel Administrativo</h4>
            </Box>

            <Box
              component="form"
              noValidate
              onSubmit={handleSubmit}
              sx={{ mt: 1 }}
            >
              <TextField
                margin="normal"
                required
                fullWidth
                id="username"
                label="Username"
                name="username"
                autoComplete="username"
                autoFocus
              />
              <TextField
                margin="normal"
                required
                fullWidth
                name="password"
                label="Password"
                type="password"
                id="password"
                autoComplete="current-password"
              />
              <Button
                type="submit"
                fullWidth
                variant="contained"
                style={{ backgroundColor: "black" }}
                sx={{ mt: 3, mb: 2 }}
              >
                {loading ? (
                  <CircularProgress color="primary" size={25} />
                ) : (
                  <Typography>Acessar</Typography>
                )}
              </Button>
            </Box>
          </Box>
        </Grid>
      </Grid>
    </ThemeProvider>
  );
}
