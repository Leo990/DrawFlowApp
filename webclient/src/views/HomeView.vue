<template>
    <div class="container">
        <h1>Programs</h1>
        <div class="row">
            <div class="col-3 py-2">
                <div class="card h-100">
                    <div class="card-body">
                        <h5 class="card-title">New Program</h5>
                        <br>
                        <RouterLink class="btn btn-success" to="/editor">
                            <span class="material-icons">
                                add
                            </span>
                        </RouterLink>
                    </div>
                </div>
            </div>
            <div class="col-3 py-2" v-for="program in programs">
                <div class="card h-100">
                    <div class="card-body">
                        <h5 class="card-title">{{ program.program_name }}</h5>
                    </div>
                    <div class="card-footer">
                        <a class="btn btn-success" @click="redirectToEditor(program.uid)">
                            <span class="material-icons">
                                play_arrow
                            </span>
                        </a>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
import { indexProgram } from '../services/apiService'
import router from '../router';
export default {
    mounted(){
        indexProgram().then((res) => {
            this.programs = res.data["programs"]
        })
    }
    ,data() {
        return {
            programs : []
        }
    },
    methods: {
        redirectToEditor(id) {
            router.push(`/editor/${id}`)
        }
    },
}
</script>
<style>
    
</style>