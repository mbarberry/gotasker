<template>
	<h1 class="text-6xl my-12"><img src="../assets/gopher-logo.png" style="display: inline-flex" width="50" height="50" />{{ msg }}</h1>
	<button class="bg-yellow-500 w-24 rounded-md my-12" type="button">Add Task</button>
	<div class="w-full h-full">
		<table class="my-0 mx-auto" v-if="processData">
			<tr>
				<th class="bg-blue-100 border px-8 py-4">Task</th>
				<th class="bg-blue-100 border px-8 py-4">Status</th>
				<th class="bg-blue-100 border text-left px-8 py-4">Mark Complete</th>
				<th class="bg-blue-100 border text-left px-8 py-4">Remove</th>
			</tr>
			<tr v-for="task in tasks" v-bind:task="task" v-bind:key="task.ID">
				<td class="border px-8 py-4">{{ task.Text }}</td>
				<td class="border px-8 py-4" :class="{ 'bg-green-300': task.Completed === 'Complete', 'bg-yellow-300': task.Completed === 'Pending' }">{{ task.Completed }}</td>
				<td class="border px-8 py-4">
					<button @click="markComplete" :disabled="task.Completed === 'Complete'">
					<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path></svg>
					</button>
				</td>
				<td class="border px-8 py-4">
					<button>
						<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
						</svg>
					</button>
				</td>
			</tr>
		</table>
	</div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },
  data () {
	return {
		tasks: [],
	}
  },
  created () {
	this.fetchData()
  }, 
  methods: {
	fetchData() {
		axios 
			.get("/api/index")
			.then(res => {
				this.tasks = res.data
			})
			.catch(err => console.err(err))
	},
	markComplete() {
		console.log('hello')
	}
  },
  computed: {
	processData: function() {
		//let processedData = []
		this.tasks.forEach(task => {
			if (task.Completed === true) {
				task.Completed = "Complete"
				// processedData.push(task)
			} else {
				task.Completed = "Pending"
				//processedData.push(task)
			}
		})
		return this.tasks//processedData
	}
  },
}
</script>

<style>
	button:disabled {
		cursor: not-allowed;
		pointer-events: all !important;
	}
</style>
