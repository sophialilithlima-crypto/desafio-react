import { useEffect, useState } from "react";
import api from "../api/api";
import "../styles/Produto.css";


function Produto(){


    const [produtos, setProdutos] = useState([]);

    const [categorias, setCategorias] = useState([]);

    const [nome, setNome] = useState("");

    const [sku, setSku] = useState("");

    const [preco, setPreco] = useState("");

    const [estoque, setEstoque] = useState("");

    const [categoriaId, setCategoriaId] = useState("");

    const [editando, setEditando] = useState(null);

    const [busca, setBusca] = useState("");

    const [pagina, setPagina] = useState(1);

    const [total, setTotal] = useState(0);

    const limite = 5;



    async function carregarProdutos(){


        try{


            const response = await api.get(
                `/produtos?page=${pagina}&limit=${limite}&q=${busca}`
            );


            setProdutos(response.data.data || []);

            setTotal(response.data.meta.total || 0);


        }catch(error){

            console.log(error);

        }


    }




    async function carregarCategorias(){


        try{


            const response = await api.get("/categorias");


            setCategorias(response.data.data || []);


        }catch(error){

            console.log(error);

        }


    }




    useEffect(()=>{


        carregarProdutos();


    },[pagina,busca]);




    useEffect(()=>{


        carregarCategorias();


    },[]);






    async function salvar(e){


        e.preventDefault();



        const produto = {


            nome,

            sku,

            preco: Number(preco),

            estoque: Number(estoque),

            categoria_id: Number(categoriaId)


        };



        try{


            if(editando){


                await api.put(
                    `/produtos/${editando}`,
                    produto
                );


            }else{


                await api.post(
                    "/produtos",
                    produto
                );


            }



            limpar();


            carregarProdutos();



        }catch(error){


            console.log(error);


        }


    }






    function editar(produto){


        setNome(produto.nome);

        setSku(produto.sku);

        setPreco(produto.preco);

        setEstoque(produto.estoque);

        setCategoriaId(produto.categoria_id);

        setEditando(produto.id);


    }






    async function excluir(id){


        const confirmar = window.confirm(
            "Deseja excluir este produto?"
        );


        if(!confirmar)
            return;



        try{


            await api.delete(
                `/produtos/${id}`
            );


            carregarProdutos();


        }catch(error){


            console.log(error);


        }


    }







    function limpar(){


        setNome("");

        setSku("");

        setPreco("");

        setEstoque("");

        setCategoriaId("");

        setEditando(null);


    }







    return(


        <div className="pagina">


            <h1>
                Produtos
            </h1>



            <input

                className="busca"

                placeholder="Buscar produto..."

                value={busca}

                onChange={(e)=>{

                    setPagina(1);

                    setBusca(e.target.value);

                }}

            />




            <div className="card">


                <form onSubmit={salvar}>


                    <input

                        placeholder="Nome"

                        value={nome}

                        onChange={
                            e=>setNome(e.target.value)
                        }

                    />



                    <input

                        placeholder="SKU"

                        value={sku}

                        onChange={
                            e=>setSku(e.target.value)
                        }

                    />



                    <input

                        type="number"

                        placeholder="Preço"

                        value={preco}

                        onChange={
                            e=>setPreco(e.target.value)
                        }

                    />



                    <input

                        type="number"

                        placeholder="Estoque"

                        value={estoque}

                        onChange={
                            e=>setEstoque(e.target.value)
                        }

                    />



                    <select

                        value={categoriaId}

                        onChange={
                            e=>setCategoriaId(e.target.value)
                        }

                    >

                        <option value="">
                            Selecione categoria
                        </option>


                        {
                            categorias.map(c=>(

                                <option 
                                    key={c.id}
                                    value={c.id}
                                >

                                    {c.nome}

                                </option>

                            ))
                        }


                    </select>



                    <button className="btn salvar">

                        {
                            editando
                            ?
                            "Atualizar"
                            :
                            "Cadastrar"
                        }

                    </button>


                </form>


            </div>





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


                    {
                        produtos.map(produto=>(


                            <tr key={produto.id}>


                                <td>
                                    {produto.id}
                                </td>


                                <td>
                                    {produto.nome}
                                </td>


                                <td>
                                    {produto.sku}
                                </td>


                                <td>
                                    R$ {produto.preco}
                                </td>


                                <td>
                                    {produto.estoque}
                                </td>


                                <td>


                                    <button
                                        className="btn editar"
                                        onClick={()=>
                                            editar(produto)
                                        }
                                    >

                                        Editar

                                    </button>


                                    <button
                                        className="btn excluir"
                                        onClick={()=>
                                            excluir(produto.id)
                                        }
                                    >

                                        Excluir

                                    </button>


                                </td>


                            </tr>


                        ))
                    }


                </tbody>


            </table>




            <div className="paginacao">


                <button

                    disabled={pagina===1}

                    onClick={()=>
                        setPagina(pagina-1)
                    }

                >

                    Anterior

                </button>



                <span>

                    Página {pagina}

                </span>




                <button

                    disabled={
                        pagina >= Math.ceil(total/limite)
                    }

                    onClick={()=>
                        setPagina(pagina+1)
                    }

                >

                    Próxima

                </button>



            </div>



        </div>


    )


}


export default Produto;