<script>
    import "../app.css"
    import Navbar from "../components/Navbar.svelte"
    export const prerender = true
</script>
<Navbar title="Tum-DEFI" />
<slot/>