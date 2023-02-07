import { useState, useEffect } from "react";
import "./App.css";

function App() {
  const [fecher, setFecher] = useState([]);

  useEffect(() => {
    try {
      (async () => {
        let res = await fetch(`http://localhost:5174/task`, {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
          },
          mode: "cors",
          cache: "default",
        });

        const data = await res.json();

        setFecher(data.data);
      })();
    } catch (error) {
      console.log(error);
    }
  });

  return (
    <div className="App">
      {fecher.map(({ description, title, done }) => {
        return (
          <>
            <h1>{title}</h1>
            <p>{description}</p>
            <p>{String(done)}</p>
          </>
        );
      })}
    </div>
  );
}

export default App;
