package timeseries

const ChartTmp = `
<html>

<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width">

<script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/3.9.1/chart.min.js" integrity="sha512-ElRFoEQdI5Ht6kZvyzXhYG9NqjtkmlkfYk0wr6wHxU9JEHakS7UJZNeml5ALk+8IKlU6jDgMabC3vkumRokgJA==" crossorigin="anonymous" referrerpolicy="no-referrer"></script>

<script>

window.addEventListener('load', (event) => {
  loadCharts()
});

const loadCharts = function() {
  loadBarChart()

  let charts = [
    {
      title: 'Jan-June',
      labels: [
        'January',
        'February',
        'March',
        'April',
        'May',
        'June'
      ],
      points: [
        0,
        10,
        5,
        2,
        20,
        30,
        45
      ],
    },
    {
      title: 'Jul-Dec',
      labels: [
        'July',
        'Aug',
        'Sep',
        'Oct',
        'Nov',
        'Dec',
      ],
      points: [
        9,
        2,
        4,
        26,
        20,
        30,
        45
      ],
    },
    {
      title: 'Mar-Aug',
      labels: [
        'Mar',
        'Apr',
        'May',
        'Jun',
        'Jul',
        'Aug',
      ],
      points: [
        87,
        56,
        40,
        26,
        20,
        30,
        45
      ]
    }
  ]

  let i = 0
  for(let chart of charts) {
    loadLineChart(i, chart)
    i++
  }
}

const loadLineChart = function(i, chart) {

  const data = {
    labels: chart.labels,
    datasets: [{
      label: chart.title || '',
      backgroundColor: 'rgb(204, 255, 153)',
      borderColor: 'rgb(204, 255, 153)',
      data: chart.points,
    }]
  };

  const config = {
    type: 'line',
    data: data,
    options: {}
  };

  let elementId = "lineChart" + i.toString()
  console.log("elementId is", elementId) 

  const myChart = new Chart(document.getElementById(elementId), config)
}

const loadBarChart = function() {
  const ctx = document.getElementById('barChart').getContext('2d');

  const myChart = new Chart(ctx, {
    type: 'bar',
    data: {
        labels: ['Red', 'Blue', 'Yellow', 'Green', 'Purple', 'Orange'],
        datasets: [{
            label: '# of Votes',
            data: [12, 19, 3, 5, 2, 3],
            backgroundColor: 'rgb(255, 99, 132)',
            borderColor: 'rgb(255, 99, 132)',
            borderWidth: 1
        }]
    },
    options: {}
  });
}
</script>
      
<style>
      html {
        font-family: sans-serif;
      }

      body {
        margin: 0;
      }

      header {
        background: purple;
        height: 100px;
      }

      h1 {
        text-align: center;
        color: white;
        line-height: 100px;
        margin: 0;
      }

      article {
        padding: 10px;
        margin: 10px;
        width: 40%;
        height: 20%;
      }

      /* Add your flexbox CSS below here */

      section {
        display: flex;
	flex-direction: row;
        flex-wrap: wrap
      }

      
</style>

</head>
<body>
    <section>
      <article>
        <h2>Test bar chart</h2>

	<canvas id="barChart"></canvas>
      </article>

      <article>
        <h2>Test line chart1</h2>

	<canvas id="lineChart0"></canvas>
      </article>

      <article>
        <h2>Test line chart2</h2>

	<canvas id="lineChart1"></canvas>
      </article>

      <article>
        <h2>Test line chart3</h2>

	<canvas id="lineChart2"></canvas>
      </article>
      <article>
        <h2>What is the story?</h2>

        <p>This is where the story explaining how the charts a related to each other goes.</p>
      </article>

    </section>
</body>
</html>
`
