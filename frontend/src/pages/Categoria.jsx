import { useEffect, useState } from "react";
import api from "../api/api";


function Categoria() {


    const [categorias, setCategorias] = useState([]);


    const [editando, setEditando] = useState(false);


    const [idEditando, setIdEditando] = useState(null);



    const [formulario, setFormulario] = useState({

        nome: ""

    });








    async function buscarCategorias() {


        try {


            const resposta = await api.get("/categorias");


            setCategorias(
                resposta.data.data || []
            );



        } catch (erro) {


            console.log(
                "Erro ao buscar categorias:",
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








    async function salvarCategoria(e) {


        e.preventDefault();



        const categoria = {


            nome: formulario.nome


        };





        try {



            if(editando){



                await api.put(

                    `/categorias/${idEditando}`,

                    categoria

                );



                alert(
                    "Categoria atualizada com sucesso!"
                );



            } else {



                await api.post(

                    "/categorias",

                    categoria

                );



                alert(
                    "Categoria cadastrada com sucesso!"
                );


            }






            limparFormulario();


            buscarCategorias();





        } catch (erro) {



            console.log(
                "Erro ao salvar categoria:",
                erro
            );



            alert(
                "Erro ao salvar categoria"
            );



        }


    }









    function editarCategoria(categoria) {



        setEditando(true);



        setIdEditando(
            categoria.id
        );



        setFormulario({

            nome: categoria.nome

        });



    }









    function limparFormulario() {



        setFormulario({

            nome: ""

        });



        setEditando(false);


        setIdEditando(null);



    }









    async function excluirCategoria(id) {



        const confirmar = window.confirm(

            "Deseja realmente excluir esta categoria?"

        );



        if(!confirmar){

            return;

        }






        try {



            await api.delete(

                `/categorias/${id}`

            );



            alert(

                "Categoria excluída com sucesso!"

            );



            buscarCategorias();





        } catch (erro) {



            console.log(

                "Erro ao excluir categoria:",

                erro

            );



            alert(

                "Erro ao excluir categoria"

            );



        }


    }









    useEffect(()=>{


        buscarCategorias();


    }, []);









    return (

        <div>



            <h1>

                Categorias

            </h1>





            <h2>

                {

                    editando

                    ?

                    "Editar Categoria"

                    :

                    "Cadastrar Categoria"

                }


            </h2>







            <form onSubmit={salvarCategoria}>



                <input


                    type="text"


                    name="nome"


                    placeholder="Nome da categoria"


                    value={formulario.nome}


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

                Categorias cadastradas

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

                            Ações

                        </th>


                    </tr>


                </thead>







                <tbody>



                    {


                        categorias.map((categoria)=>(



                            <tr key={categoria.id}>


                                <td>

                                    {categoria.id}

                                </td>





                                <td>

                                    {categoria.nome}

                                </td>





                                <td>



                                    <button


                                        onClick={() =>
                                            editarCategoria(categoria)
                                        }


                                    >

                                        Editar


                                    </button>







                                    <button


                                        onClick={() =>
                                            excluirCategoria(categoria.id)
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



export default Categoria;