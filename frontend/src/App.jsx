import { useEffect, useState } from "react"

function App() {

  const [mensagem, setMensagem] = useState("")

  useEffect(() => {

    fetch("http://localhost:8080")
      .then(res => res.text())
      .then(data => setMensagem(data))

  }, [])


  return (
    <>
      <h1>Frontend React</h1>
      <p>{mensagem}</p>
    </>
  )
}

export default App