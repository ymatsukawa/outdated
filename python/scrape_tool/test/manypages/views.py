from django.shortcuts import render
from django.core.paginator import Paginator, EmptyPage, PageNotAnInteger
import copy

_MY_FOODS = ['apple', 'banana', 'orrange', 'grape', 'spam']
_ITEM_IN_EACH_PAGE = 2

def index(request):
    foods = copy.deepcopy(_MY_FOODS)

    request_page = request.GET.get('page')
    paginator = Paginator(foods, _ITEM_IN_EACH_PAGE)

    try:
        pages = paginator.get_page(request_page)
    except PageNotAnInteger:
        pages = paginator.get_page(1)
    except EmptyPage:
        pages = paginator.get_page(paginator.num_pages)

    return render(request, 'manypages/index.html', { 'pages': pages })