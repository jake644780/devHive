import { useState } from "react";
import styles from "./style.module.scss";

export default function AddIdea(){
    let [category, setCategory] = useState("");
    let [content, setContent] = useState("");
    let [error, setError] = useState("");


    const handleCategoryChange = (event) => {
        setCategory(event.target.value); // Update category state
    };

    const handleContentChange = (event) => {
        setContent(event.target.value); // Update content state
    };

    const handleSubmit = (event) => {
        event.preventDefault();
        console.log({ category, content });
        let data = {category, content};
        const fetchData = async () => {
            try {
              const response = await fetch('http://localhost:8080/idea/new',{
                method: 'POST',
                headers: {
                  'Content-Type': 'application/json',
                },
                body: JSON.stringify(data),
              });
              const result = await response.json();
              setCategory("");
              setContent("");
              alert("idea added");
            } catch (error) {
              setError(error);
            }
          };
          fetchData();
    };

    
    return (
    <div className={styles.content}>
        <form className={styles.container}>
            <div className={styles.title}>submit a new idea</div>
            <label htmlFor="category">category</label>
            <input type="text" id="category" placeholder="microservices, ml, etc." onChange={handleCategoryChange} value={category} />
            <label htmlFor="content">content</label>
            <textarea id="content" onChange={handleContentChange} value={content}></textarea>
            <button type="submit" onClick={handleSubmit}>submit</button>
        </form>
    </div>
    )
}