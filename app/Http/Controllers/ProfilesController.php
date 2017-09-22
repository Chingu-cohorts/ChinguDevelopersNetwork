<?php

namespace App\Http\Controllers;

use App\User;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Validator;

class ProfilesController extends Controller
{
	protected function validator(array $data)
	{
		return Validator::make($data, [
			'location' => 'string|min:3|max:20',
			'about' => 'string|max:255',
			'github' => 'string|min:3|max:30',
			'medium' => 'string|min:3|max:30',
			'twitter' => 'string|min:3|max:30',
		]);
	}

    public function show($slug)
    {
    	$user = User::where('slug', $slug)->first();
    	return view('profiles.show')->with('user', $user);
    }
}
