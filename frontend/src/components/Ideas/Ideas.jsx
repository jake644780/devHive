import React, { useState, useEffect } from 'react';
import styles from "./style.module.scss";

export default function Ideas (){
    const [ideas, setIdeas] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
  

    const deleteIdeaFromBackend = async (index) => {
      try {
        const response = await fetch(`http://localhost:8080/idea/${index}`, {method: "DELETE"});
        const result = await response.json();
      } catch (error) {
        setError(error);
      }
    };

    const deleteIdea = (index) => {
      const updatedIdeas = ideas.filter((idea) => idea.id !== index);
      setIdeas(updatedIdeas);
      deleteIdeaFromBackend(index);
    };

    const fetchIdeas = async () => {
      try {
        setLoading(true);
        const response = await fetch('http://localhost:8080/ideas');
        const result = await response.json();
        
        setIdeas(result);
      } catch (error) {
        setError(error);
      } finally {
        setLoading(false);
      }
    };

    useEffect(() => {
        fetchIdeas();
      }, []);
    

      

      if (loading) return <p>Loading...</p>;
      if (error) return <p>Error: {error.message}</p>;
      


      return (
        <div>
          <h1>Fetched Data</h1>
            {
            ideas.map((idea, i) => (
              <div key={idea.id}>
                <h3>Idea {idea.id}:</h3>
                <p><strong>Category:</strong> {idea.category}</p>
                <p><strong>Content:</strong> {idea.content}</p>
                <button onClick={()=>deleteIdea(idea.id)}>delete idea</button>
              </div>
            ))
            }
        </div>
      );


};3