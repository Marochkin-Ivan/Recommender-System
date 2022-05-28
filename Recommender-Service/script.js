
async function ww() {
    let response = await fetch("127.0.0.1");

    if (response.ok) { // если HTTP-статус в диапазоне 200-299
        // получаем тело ответа (см. про этот метод ниже)
        let json = await response.json();
    } else {
        alert("Ошибка HTTP: " + response.status);
    }
    return response.status
}




console.log(ww())