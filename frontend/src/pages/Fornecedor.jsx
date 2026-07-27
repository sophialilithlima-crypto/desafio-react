import { useEffect, useState } from "react";
import api from "../api/api";


function Fornecedor() {


    const [fornecedores, setFornecedores] = useState([]);


    const [editando, setEditando] = useState(false);


    const [idEditando, setIdEditando] = useState(null);



    const [formulario, setFormulario] = useState({

        nome: "",
        email: "",
        telefone: ""

    });







    async function buscarFornecedores() {


        try {


            const resposta = await api.get("/fornecedores");


            setFornecedores(
                resposta.data.data || []
            );



        } catch (erro) {


            console.log(
                "Erro ao buscar fornecedores:",
                erro
            );


        }


    }








    function alterarCampo(e) {


        setFormulario({

            ...formulario,

            [e.target.name]: e.target.value

        });


    }









    function validarEmail(email) {


        return email.includes("@");

    }









    async function salvarFornecedor(e) {


        e.preventDefault();




        if(!validarEmail(formulario.email)){


            alert(
                "Digite um email válido"
            );


            return;


        }






        const fornecedor = {


            nome: formulario.nome,

            email: formulario.email,

            telefone: formulario.telefone


        };








        try {




            if(editando){





                await api.put(

                    `/fornecedores/${idEditando}`,

                    fornecedor

                );



                alert(

                    "Fornecedor atualizado com sucesso!"

                );





            } else {





                await api.post(

                    "/fornecedores",

                    fornecedor

                );



                alert(

                    "Fornecedor cadastrado com sucesso!"

                );



            }






            limparFormulario();


            buscarFornecedores();






        } catch (erro) {



            console.log(

                "Erro ao salvar fornecedor:",

                erro

            );



            alert(

                "Erro ao salvar fornecedor"

            );



        }


    }









    function editarFornecedor(fornecedor) {



        setEditando(true);



        setIdEditando(

            fornecedor.id

        );



        setFormulario({


            nome: fornecedor.nome,


            email: fornecedor.email,


            telefone: fornecedor.telefone


        });



    }









    function limparFormulario(){



        setFormulario({


            nome: "",


            email: "",


            telefone: ""


        });




        setEditando(false);



        setIdEditando(null);



    }









    async function excluirFornecedor(id){



        const confirmar = window.confirm(

            "Deseja realmente excluir este fornecedor?"

        );




        if(!confirmar){

            return;

        }








        try {




            await api.delete(

                `/fornecedores/${id}`

            );




            alert(

                "Fornecedor excluído com sucesso!"

            );




            buscarFornecedores();






        } catch (erro) {



            console.log(

                "Erro ao excluir fornecedor:",

                erro

            );



            alert(

                "Erro ao excluir fornecedor"

            );



        }


    }









    useEffect(()=>{


        buscarFornecedores();


    }, []);









    return (

        <div>



            <h1>

                Fornecedores

            </h1>







            <h2>


                {

                    editando

                    ?

                    "Editar Fornecedor"

                    :

                    "Cadastrar Fornecedor"

                }


            </h2>







            <form onSubmit={salvarFornecedor}>


                <input

                    type="text"

                    name="nome"

                    placeholder="Nome"

                    value={formulario.nome}

                    onChange={alterarCampo}

                />





                <input

                    type="email"

                    name="email"

                    placeholder="Email"

                    value={formulario.email}

                    onChange={alterarCampo}

                />





                <input

                    type="text"

                    name="telefone"

                    placeholder="Telefone"

                    value={formulario.telefone}

                    onChange={alterarCampo}

                />







                <button type="submit">


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

                        onClick={limparFormulario}

                    >

                        Cancelar

                    </button>


                }





            </form>









            <h2>

                Fornecedores cadastrados

            </h2>









            <table border="1">


                <thead>


                    <tr>


                        <th>

                            ID

                        </th>


                        <th>

                            Nome

                        </th>


                        <th>

                            Email

                        </th>


                        <th>

                            Telefone

                        </th>


                        <th>

                            Ações

                        </th>


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

                                        onClick={() =>
                                            editarFornecedor(fornecedor)
                                        }

                                    >

                                        Editar

                                    </button>






                                    <button

                                        onClick={() =>
                                            excluirFornecedor(fornecedor.id)
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


    );


}



export default Fornecedor;