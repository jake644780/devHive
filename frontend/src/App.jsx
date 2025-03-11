import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './css/App.css'

function App() {
  const [count, setCount] = useState(0);

  const [likes, setLikes] = useState(10);

  return (
    <>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <textarea onInput={()=>setLikes((likes)=> likes - 1)}>

        </textarea>
        <div>
          {likes}
        </div>
      </div>
    </>
  )
}

export default App;