document.addEventListener("DOMContentLoaded", async () => {
    const response = await fetch('http://localhost:8080/events');
    const events = await response.json();
    const eventsContainer = document.getElementById('events');

    events.forEach(event => {
        const eventElement = document.createElement('div');
        eventElement.innerHTML = `
            <h2>${event.name}</h2>
            <p>${event.description}</p>
            <p>${event.date} ${event.time} at ${event.location}</p>
            <p>Organized by: ${event.organizer}</p>
        `;
        eventsContainer.appendChild(eventElement);
    });
});
