import { useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "../api/api";
import "../styles/Login.css";

function Login() {
    const navigate = useNavigate();
    const [usuario, setUsuario] = useState("");
    const [senha, setSenha] = useState("");
    const [erro, setErro] = useState("");
    const [carregando, setCarregando] = useState(false);

    async function entrar(e) {
        e.preventDefault();
        setErro("");
        setCarregando(true);

        try {
            const response = await api.post("/auth/login", {
                usuario,
                senha
            });

            localStorage.setItem("token", response.data.data.token);
            navigate("/");
        } catch (error) {
            setErro(error.response?.data?.error || "Não foi possível entrar");
        } finally {
            setCarregando(false);
        }
    }

    return (
        <div className="login-page">
            <div className="login-card">
                <h1>Login</h1>

                <form onSubmit={entrar}>
                    <input
                        placeholder="Usuário"
                        value={usuario}
                        onChange={(e) => setUsuario(e.target.value)}
                    />

                    <input
                        type="password"
                        placeholder="Senha"
                        value={senha}
                        onChange={(e) => setSenha(e.target.value)}
                    />

                    <button type="submit" disabled={carregando}>
                        {carregando ? "Entrando..." : "Entrar"}
                    </button>
                </form>

                {erro && <p className="login-erro">{erro}</p>}
            </div>
        </div>
    );
}

export default Login;
