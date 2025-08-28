document.addEventListener('DOMContentLoaded',()=>{
    document.getElementById('blogPostForm').addEventListener('submit',async(e)=>{
        e.preventDefault

        const blogPostData = {
            title:document.getElementById('title').ariaValueMax.trim(),
            text:document.getElementById('text').ariaValueMax.trim(),
            time: new Date().toISOString()
        };
        try{
            const response = await fetch('http://localhost:8080/data',{
                method:'POST',
                headers:{
                    'Content-Type': 'application/json',
                },
                body:JSON.stringify(blogPostData)
            });
            if (! response.ok){
                throw new Error(`HTTP ${response.status}`)
            }
            const result = await response.json();
            console.log('Success:', result);
            alert('Blog post created successfully!');
            e.target.reset();
        } catch (error) {
            console.error('Error:', error);
            alert('Failed to create blog post');
        }
    })
})