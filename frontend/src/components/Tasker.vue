<template>
  <div>
    <h1 class="text-6xl my-12">
      <img src="../assets/gopher-logo.png" style="display: inline-flex" width="50" height="50" />{{ msg }}
    </h1>
    <div>
      <label class="uppercase italic" for="task">New Task: </label>
      <input class="bg-gray-200" v-model="task" />
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
          <td class="table-data">
            {{ task.Text }}
          </td>
          <td class="table-data" :class="{
            'bg-green-300': task.Completed,
            'bg-yellow-300': !task.Completed,
          }">
            {{ task.Completed ? 'Complete' : 'Pending' }}
          </td>
          <td class="table-data">
            <button @click="markComplete(task.ID)" :disabled="task.Completed" class="disabled:cursor-not-allowed">
              <svg xmlns="http://www.w3.org/2000/svg" fill="currentColor" class="w-6 h-6" viewBox="0 0 448 512">

                <path
                  d="M438.6 105.4c12.5 12.5 12.5 32.8 0 45.3l-256 256c-12.5 12.5-32.8 12.5-45.3 0l-128-128c-12.5-12.5-12.5-32.8 0-45.3s32.8-12.5 45.3 0L160 338.7 393.4 105.4c12.5-12.5 32.8-12.5 45.3 0z" />
              </svg>
            </button>
          </td>
          <td class="table-data">
            <button @click="deleteTask(task.ID)">
              <svg class="w-6 h-6" fill="white" stroke="currentColor" viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16">
                </path>
              </svg>
            </button>
          </td>
        </tr>
      </table>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

import { API_URL } from '../constants';

export default {
  name: 'Tasker',
  props: {
    msg: String,
  },
  data() {
    return {
      task: '',
      tasks: [],
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      axios
        .get(
          `${API_URL}/getTasks`
        )
        .then((res) => {
          this.tasks = res.data;
        })
        .catch((err) => console.error(err));
    },
    handleSubmit() {
      axios
        .post(
          `${API_URL}/addTask`,
          { task: this.task }
        )
        .then((res) => {
          const { data: { Text, Completed, ID, UpdatedAt, CreatedAt } } = res;
          this.tasks.push({
            Text, Completed, ID, UpdatedAt, CreatedAt
          });
          this.task = '';
        })
        .catch((err) => console.error(err));
    },
    markComplete(id) {
      axios
        .post(
          `${API_URL}/updateStatus`,
          { task: id }
        )
        .then((res) => {
          const task = this.tasks.find((task) => task.ID === res.data);
          task.Completed = true;
        })
        .catch((err) => console.error(err));
    },
    deleteTask(id) {
      axios
        .post(
          `${API_URL}/deleteTask`,
          { task: id }
        )
        .then((res) => {
          const taskIndex = this.tasks.findIndex((task) => task.ID === res.data);
          const updatedTasks = [
            ...this.tasks.slice(0, taskIndex),
            ...this.tasks.slice(taskIndex + 1, this.tasks.length),
          ];
          this.tasks = updatedTasks;
        })
        .catch((err) => console.error(err));
    },
  },
};
</script>
