<script setup>
import axios from "axios"
import configFile from '../../config.json'
import { onMounted, ref } from "vue"

const columns = ref([])
const showNewTaskForm = ref(false)
const newTask = ref({
  name: "",
  status: "TO DO",
})

const loadColumns = async () => {
  const response = await axios.get(configFile.API_BASE_URL + "/tasks")
  columns.value = response.data.tasks
}

const openNewTaskForm = () => {
  showNewTaskForm.value = true
}

const closeNewTaskForm = () => {
  showNewTaskForm.value = false
}

const submitNewTask = async () => {
  await axios.post(configFile.API_BASE_URL + "/tasks", newTask.value)
  newTask.value = { name: "", status: "TO DO" }
  showNewTaskForm.value = false
  await loadColumns()
}

onMounted(async () => {
  await loadColumns()
})
</script>

<template>
  <section class="board">
    <header class="board-header">
      <div>
        <p class="eyebrow">Dashboard</p>
        <h2>Tasks board</h2>
        <p class="muted">Track the status of active workstreams.</p>
      </div>
      <div class="board-actions">
        <button class="btn ghost">Filter</button>
        <button class="btn primary" @click="openNewTaskForm">New task</button>
      </div>
    </header>

    <div v-if="showNewTaskForm" class="task-form-card">
      <div class="task-form-header">
        <div>
          <p class="eyebrow">New task</p>
          <h3>Create task</h3>
          <p class="muted">Add a task to your board.</p>
        </div>
        <button class="btn ghost" type="button" @click="closeNewTaskForm">Close</button>
      </div>
      <form class="task-form" @submit.prevent="submitNewTask">
        <label class="field">
          <span>Name</span>
          <input v-model="newTask.name" type="text" placeholder="Task title" required />
        </label>
        <label class="field">
          <span>Status</span>
          <select v-model="newTask.status">
            <option value="TO DO">TO DO</option>
            <option value="IN PROGRESS">IN PROGRESS</option>
            <option value="DONE">DONE</option>
          </select>
        </label>
        <div class="task-form-actions">
          <button class="btn ghost" type="button" @click="closeNewTaskForm">Cancel</button>
          <button class="btn primary" type="submit">Create task</button>
        </div>
      </form>
    </div>

    <div class="board-columns">
      <div v-for="column in columns" :key="column.Id" class="column">
        <div class="column-header">
          <h3>{{ column.Name }}</h3>
          <span class="pill neutral">{{ column.Status }}</span>
        </div>
        
      </div>
    </div>
  </section>
</template>
