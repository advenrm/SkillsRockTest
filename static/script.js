// function for message box when deleting
function confirmDelete(id, title) {
    Swal.fire({
        title: `Are you sure delete ${title}?`, 
        text: 'You will not be able to recover this item!',
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#d33',
        cancelButtonColor: '#3085d6',
        confirmButtonText: 'Yes, delete it!'
    }).then((result) => {
        if (result.isConfirmed) {
            document.getElementById('deleteForm' + id).submit(); 
        }
    });
    return false;
}

$(document).ready(function() {
    $('#dataTable').DataTable({
        "searching": true,
    });
});