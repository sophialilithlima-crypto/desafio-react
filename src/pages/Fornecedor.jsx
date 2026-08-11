import { useEffect, useState } from "react";
import api from "../api/api";
import "../styles/Fornecedor.css";


function Fornecedor() {


    const [fornecedores, setFornecedores] = useState([]);

    const [nome, setNome] = useState("");

    const [email, setEmail] = useState("");

    const [telefone, setTelefone] = useState("");

    const [editando, setEditando] = useState(null);

    const [mensagem, setMensagem] = useState("");




    async function carregarFornecedores(){

        try{

            const response = await api.get("/fornecedores");

            setFornecedores(response.data.data || []);

        }catch(error){

            console.log(error);

        }

    }





    useEffect(()=>{

        carregarFornecedores();

    },[]);







    async function salvar(e){

        e.preventDefault();


        if(!nome.trim()){

            setMensagem("Informe o nome do fornecedor");

            return;

        }



        try{


            const dados = {

                nome,
                email,
                telefone

            };



            if(editando){


                await api.put(
                    `/fornecedores/${editando}`,
                    dados
                );


                setMensagem(
                    "Fornecedor atualizado com sucesso"
                );


            }else{


                await api.post(
                    "/fornecedores",
                    dados
                );


                setMensagem(
                    "Fornecedor cadastrado com sucesso"
                );

            }



            limpar();

            carregarFornecedores();



        }catch(error){

            console.log(error);

            setMensagem(
                "Erro ao salvar fornecedor"
            );

        }


    }







    function editar(fornecedor){


        setNome(fornecedor.nome);

        setEmail(fornecedor.email);

        setTelefone(fornecedor.telefone);

        setEditando(fornecedor.id);

        setMensagem("");

    }







    async function excluir(id){


        const confirmar = window.confirm(
            "Deseja excluir este fornecedor?"
        );


        if(!confirmar){

            return;

        }



        try{


            await api.delete(
                `/fornecedores/${id}`
            );


            setMensagem(
                "Fornecedor removido"
            );


            carregarFornecedores();



        }catch(error){

            console.log(error);

            setMensagem(
                "Erro ao excluir fornecedor"
            );

        }


    }






    function limpar(){


        setNome("");

        setEmail("");

        setTelefone("");

        setEditando(null);


    }








    return(


        <div className="pagina">


            <h1>
                Fornecedores
            </h1>



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
                        placeholder="Email"
                        value={email}
                        onChange={
                            e=>setEmail(e.target.value)
                        }
                    />



                    <input
                        placeholder="Telefone"
                        value={telefone}
                        onChange={
                            e=>setTelefone(e.target.value)
                        }
                    />




                    <button className="btn salvar">

                        {
                            editando
                            ?
                            "Atualizar"
                            :
                            "Cadastrar"
                        }

                    </button>



                    {
                        editando &&

                        <button
                            type="button"
                            className="btn cancelar"
                            onClick={limpar}
                        >
                            Cancelar
                        </button>
                    }



                </form>



                {
                    mensagem &&
                    <p className="mensagem">
                        {mensagem}
                    </p>
                }


            </div>





            <div className="tabela-container">


                <table>


                    <thead>

                        <tr>

                            <th>ID</th>

                            <th>Nome</th>

                            <th>Email</th>

                            <th>Telefone</th>

                            <th>Ações</th>

                        </tr>

                    </thead>




                    <tbody>


                        {
                            fornecedores.map((fornecedor)=>(


                                <tr key={fornecedor.id}>


                                    <td>
                                        {fornecedor.id}
                                    </td>


                                    <td>
                                        {fornecedor.nome}
                                    </td>


                                    <td>
                                        {fornecedor.email}
                                    </td>


                                    <td>
                                        {fornecedor.telefone}
                                    </td>


                                    <td>


                                        <button
                                            className="btn editar"
                                            onClick={
                                                ()=>editar(fornecedor)
                                            }
                                        >
                                            Editar
                                        </button>



                                        <button
                                            className="btn excluir"
                                            onClick={
                                                ()=>excluir(fornecedor.id)
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


            </div>



        </div>


    )


}


export default Fornecedor;