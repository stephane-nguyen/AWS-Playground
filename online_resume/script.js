// Sample data for employment history
const employmentData = [
  {
    company: "SAP",
    image: "image/sap-logo.png",
    position: "Full Stack Developer Apprentice",
    location: "Levallois-Perret, Ile-de-France, France",
    years: "March 2022 - September 2024",
  },
  {
    company: "SAP",
    image: "image/sap-logo.png",
    position: "Software Testing and Automation Intern",
    location: "Palo Alto, California, United States",
    years: "June 2023 - September 2023",
  },
  {
    company: "Thales",
    image: "image/Thales-logo.png",
    position: "Applied Mathematics Apprentice",
    location: "Gennevilliers, Ile-de-France, France",
    years: "September 2021 - November 2021",
  },
];

// Function to populate the employment timeline dynamically
function populateEmploymentTimeline() {
  const timeline = document.getElementById("employment-timeline");
  employmentData.forEach((job) => {
    const event = document.createElement("div");
    event.classList.add("employment-item");

    event.innerHTML = `
    <img src="${job.image}" alt="${job.company}" class="company-logo">
    <span class="position"><strong>${job.position}</strong></span> <br />
    <span class="location">${job.location}</span> <br />
    <span class="years">${job.years}</span> <br />
    `;

    timeline.appendChild(event);
  });
}

// Function to populate introduction content dynamically
function populateIntroduction() {
  const introductionContent = document.getElementById("introduction-content");
  introductionContent.innerHTML = `
          <p>
              Stéphane, a 22-year-old pursuing the last year of a master's degree in Computer Science
              in France. Currently, I am also gaining practical experience as an apprentice at SAP
              France where I actively contribute to the development of internal tools aimed at
              facilitating stakeholders’ workflow. I do enjoy working in a team and learning from
              people. I believe everyone has something new and unique to bring to the table. I find
              great satisfaction in the process of building products that help people through
              problem-solving. That is why I chose the path of Software Engineering. I consider myself
              as humble, curious, and resilient. I always seek feedback to keep improving myself. I
              have had the opportunity to complete an internship in California, United States, I
              discovered a profound appreciation for diverse cultures and work environments. It is
              this exposure that motivates me to seek further experiences outside my comfort zone.
          </p>
      `;
}

// Populate employment timeline and introduction content on page load
window.onload = function () {
  populateEmploymentTimeline();
  populateIntroduction();
};
