<template>
    <div>
        <h1 class="text-6xl my-12"><img src="../assets/gopher-logo.png" style="display: inline-flex" width="50" height="50" />{{ msg }}</h1>
        <div>
            <label class="uppercase italic" for="task">New Task: </label>
            <input class="bg-gray-200" v-model="task">
            <button @click="handleSubmit" class="bg-blue-400 w-24 my-12">Add</button>
        </div>
        <div class="w-full h-full">
            <table class="my-0 mx-auto">
                <tr>
                    <th class="table-header">Task</th>
                    <th class="table-header">Status</th>
                    <th class="table-header text-left">Mark Complete</th>
                    <th class="table-header text-left">Remove</th>
                </tr>
                <tr v-for="task in tasks" v-bind:task="task" v-bind:key="task.ID">
                    <td class="table-data">{{ task.Text }}</td>
                    <td class="table-data" :class="{ 'bg-green-300': task.Status === 'Complete', 'bg-yellow-300': task.Status === 'Pending' }">{{ task.Status }}</td>
                    <td class="table-data">
                        <button class="disabled:cursor-not-allowed" @click="markComplete(task.Text)" :disabled="task.Status === 'Complete'">
                        <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path></svg>
                        </button>
                    </td>
                    <td class="table-data">
                        <button @click="deleteTask(task.Text)">
                            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                            </svg>
                        </button>
                    </td>
                </tr>
            </table>
        </div>
    </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Tasker',
  props: {
    msg: String
  },
  data () {
	return {
		task: "",
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
				res.data.forEach(task => {
					(task.Completed === true) ? task.Status = "Complete" : task.Status = "Pending"
				})
				this.tasks = res.data
			})	
			.catch(err => console.error(err))
	},
	handleSubmit() {
		const refreshData = this.fetchData
		axios
			.post("/api/add", {task: this.task})
			.then(() => {
				refreshData()
				this.task = ""
			})
			.then(() => console.log("Task successfully added."))
			.catch(err => console.error(err))
	},
	markComplete(text) {
		const refreshData = this.fetchData
		axios
			.post("/api/done", {task: text})
			.then(() => refreshData())
			.then(() => console.log("Task successfully marked complete."))
			.catch(err => console.error(err))
	},
	deleteTask(text) {
		const refreshData = this.fetchData
		axios
			.post("/api/rm", {task: text})
			.then(() => refreshData())
			.then(() => console.log("Task successfully deleted."))
			.catch(err => console.error(err))
	}
  },
}
</script>
