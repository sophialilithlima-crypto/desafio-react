import { useEffect, useState } from "react";
import api from "../api/api";
import "../styles/Produto.css";

function Produto() {
    const [produtos, setProdutos] = useState([]);
    const [categorias, setCategorias] = useState([]);
    const [fornecedores, setFornecedores] = useState([]);

    const [nome, setNome] = useState("");
    const [sku, setSku] = useState("");
    const [preco, setPreco] = useState("");
    const [estoque, setEstoque] = useState("");
    const [categoriaId, setCategoriaId] = useState("");
    const [fornecedorIds, setFornecedorIds] = useState([]);

    const [editando, setEditando] = useState(null);
    const [busca, setBusca] = useState("");
    const [pagina, setPagina] = useState(1);
    const [total, setTotal] = useState(0);

    const [carregando, setCarregando] = useState(false);
    const [mensagem, setMensagem] = useState("");
    const limite = 5;

    async function carregarProdutos() {
        setCarregando(true);

        try {
            const response = await api.get(
                `/produtos?page=${pagina}&limit=${limite}&q=${encodeURIComponent(busca)}`
            );

            setProdutos(response.data.data || []);
            setTotal(response.data.meta?.total || 0);
        } catch (error) {
            setMensagem(error.response?.data?.error || "Erro ao carregar produtos");
        } finally {
            setCarregando(false);
        }
    }

    async function carregarRelacionamentos() {
        try {
            const [categoriasResponse, fornecedoresResponse] = await Promise.all([
                api.get("/categorias?limit=100"),
                api.get("/fornecedores?limit=100")
            ]);

            setCategorias(categoriasResponse.data.data || []);
            setFornecedores(fornecedoresResponse.data.data || []);
        } catch (error) {
            setMensagem("Erro ao carregar categorias ou fornecedores");
        }
    }

    useEffect(() => {
        carregarProdutos();
    }, [pagina, busca]);

    useEffect(() => {
        carregarRelacionamentos();
    }, []);

    function alternarFornecedor(id) {
        setFornecedorIds((atuais) =>
            atuais.includes(id)
                ? atuais.filter((item) => item !== id)
                : [...atuais, id]
        );
    }

    async function salvar(e) {
        e.preventDefault();
        setMensagem("");

        const produto = {
            nome,
            sku,
            preco: Number(preco),
            estoque: Number(estoque),
            categoria_id: Number(categoriaId),
            fornecedor_ids: fornecedorIds
        };

        try {
            if (editando) {
                await api.put(`/produtos/${editando}`, produto);
                setMensagem("Produto atualizado com sucesso");
            } else {
                await api.post("/produtos", produto);
                setMensagem("Produto cadastrado com sucesso");
            }

            limpar();
            await carregarProdutos();
        } catch (error) {
            setMensagem(error.response?.data?.error || "Erro ao salvar produto");
        }
    }

    function editar(produto) {
        setNome(produto.nome);
        setSku(produto.sku);
        setPreco(produto.preco);
        setEstoque(produto.estoque);
        setCategoriaId(produto.categoria_id);
        setFornecedorIds(produto.fornecedor_ids || []);
        setEditando(produto.id);
        setMensagem("");
    }

    async function excluir(id) {
        if (!window.confirm("Deseja excluir este produto?")) {
            return;
        }

        try {
            await api.delete(`/produtos/${id}`);
            setMensagem("Produto removido com sucesso");
            await carregarProdutos();
        } catch (error) {
            setMensagem(error.response?.data?.error || "Erro ao excluir produto");
        }
    }

    function limpar() {
        setNome("");
        setSku("");
        setPreco("");
        setEstoque("");
        setCategoriaId("");
        setFornecedorIds([]);
        setEditando(null);
    }

    return (
        <div className="pagina">
            <h1>Produtos</h1>

            <input
                className="busca"
                placeholder="Buscar produto..."
                value={busca}
                onChange={(e) => {
                    setPagina(1);
                    setBusca(e.target.value);
                }}
            />

            <div className="card">
                <form onSubmit={salvar}>
                    <input
                        placeholder="Nome"
                        value={nome}
                        onChange={(e) => setNome(e.target.value)}
                    />

                    <input
                        placeholder="SKU"
                        value={sku}
                        onChange={(e) => setSku(e.target.value)}
                    />

                    <input
                        type="number"
                        step="0.01"
                        placeholder="Preço"
                        value={preco}
                        onChange={(e) => setPreco(e.target.value)}
                    />

                    <input
                        type="number"
                        placeholder="Estoque"
                        value={estoque}
                        onChange={(e) => setEstoque(e.target.value)}
                    />

                    <select
                        value={categoriaId}
                        onChange={(e) => setCategoriaId(e.target.value)}
                    >
                        <option value="">Selecione categoria</option>

                        {categorias.map((categoria) => (
                            <option key={categoria.id} value={categoria.id}>
                                {categoria.nome}
                            </option>
                        ))}
                    </select>

                    <div className="fornecedores">
                        <strong>Fornecedores</strong>

                        {fornecedores.map((fornecedor) => (
                            <label key={fornecedor.id}>
                                <input
                                    type="checkbox"
                                    checked={fornecedorIds.includes(fornecedor.id)}
                                    onChange={() => alternarFornecedor(fornecedor.id)}
                                />
                                {fornecedor.nome}
                            </label>
                        ))}
                    </div>

                    <div>
                        <button className="btn salvar">
                            {editando ? "Atualizar" : "Cadastrar"}
                        </button>

                        {editando && (
                            <button
                                type="button"
                                className="btn cancelar"
                                onClick={limpar}
                            >
                                Cancelar
                            </button>
                        )}
                    </div>
                </form>

                {mensagem && <p className="mensagem">{mensagem}</p>}
            </div>

            {carregando ? (
                <p>Carregando...</p>
            ) : produtos.length === 0 ? (
                <p>Nenhum produto encontrado.</p>
            ) : (
                <table>
                    <thead>
                        <tr>
                            <th>ID</th>
                            <th>Nome</th>
                            <th>SKU</th>
                            <th>Preço</th>
                            <th>Estoque</th>
                            <th>Ações</th>
                        </tr>
                    </thead>

                    <tbody>
                        {produtos.map((produto) => (
                            <tr key={produto.id}>
                                <td>{produto.id}</td>
                                <td>{produto.nome}</td>
                                <td>{produto.sku}</td>
                                <td>R$ {produto.preco}</td>
                                <td>{produto.estoque}</td>
                                <td>
                                    <button
                                        className="btn editar"
                                        onClick={() => editar(produto)}
                                    >
                                        Editar
                                    </button>

                                    <button
                                        className="btn excluir"
                                        onClick={() => excluir(produto.id)}
                                    >
                                        Excluir
                                    </button>
                                </td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            )}

            <div className="paginacao">
                <button
                    disabled={pagina === 1}
                    onClick={() => setPagina(pagina - 1)}
                >
                    Anterior
                </button>

                <span>
                    Página {pagina}
                </span>

                <button
                    disabled={pagina >= Math.ceil(total / limite)}
                    onClick={() => setPagina(pagina + 1)}
                >
                    Próxima
                </button>
            </div>
        </div>
    );
}

export default Produto;
