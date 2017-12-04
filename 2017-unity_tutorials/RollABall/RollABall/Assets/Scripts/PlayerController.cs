using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class PlayerController : MonoBehaviour {

	public int speed;
	public Text scoreText;
	public Text winText;

	private int score;

	private Rigidbody rigidBody;

	void Start () {
		score = 0;
		UpdateScoreText ();
		winText.gameObject.SetActive (false);
		rigidBody = GetComponent<Rigidbody> ();
	}

	void FixedUpdate () {
		float horizontalAxis = Input.GetAxis ("Horizontal");
		float verticalAxis = Input.GetAxis ("Vertical");

		rigidBody.AddForce (new Vector3 (horizontalAxis, 0.0f, verticalAxis) * speed);
	}

	void OnTriggerEnter(Collider other) {
		if (other.CompareTag("PickUp")) {
			other.gameObject.SetActive (false);
			score += 1;
			UpdateScoreText ();
		}

		if (score == 10) {
			gameObject.SetActive (false);
			winText.gameObject.SetActive (true);
		}
	}

	private void UpdateScoreText () {
		scoreText.text = "Score: " + score.ToString();
	}
}
